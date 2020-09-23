package jobserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/api/present"
	"github.com/concourse/concourse/atc/db"
)

func (s *Server) ListJobBuilds(pipeline db.Pipeline) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			builds     []db.Build
			pagination db.Pagination
			err        error
			from       int
			to         int
			limit      int
		)

		logger := s.logger.Session("list-job-builds")

		jobName := r.FormValue(":job_name")

		timestamps := r.FormValue(atc.PaginationQueryTimestamps)

		urlFrom := r.FormValue(atc.PaginationQueryFrom)
		urlTo := r.FormValue(atc.PaginationQueryTo)

		urlLimit := r.FormValue(atc.PaginationQueryLimit)
		limit, _ = strconv.Atoi(urlLimit)
		if limit == 0 {
			limit = atc.PaginationAPIDefaultLimit
		}

		page := db.Page{Limit: limit}
		if urlFrom != "" {
			from, _ = strconv.Atoi(urlFrom)
			page.From = db.NewIntPtr(from)
		}
		if urlTo != "" {
			to, _ = strconv.Atoi(urlTo)
			page.To = db.NewIntPtr(to)
		}

		job, found, err := pipeline.Job(jobName)
		if err != nil {
			logger.Error("failed-to-get-job", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if timestamps == "" {
			builds, pagination, err = job.Builds(page)
		} else {
			page.UseDate = true
			builds, pagination, err = job.BuildsWithTime(page)
		}
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if pagination.Older != nil {
			s.addNextLink(w, pipeline.ID(), jobName, *pagination.Older)
		}

		if pagination.Newer != nil {
			s.addPreviousLink(w, pipeline.ID(), jobName, *pagination.Newer)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jobBuilds := make([]atc.Build, len(builds))
		for i := 0; i < len(builds); i++ {
			jobBuilds[i] = present.Build(builds[i])
		}

		err = json.NewEncoder(w).Encode(jobBuilds)
		if err != nil {
			logger.Error("failed-to-encode-job-builds", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

func (s *Server) addNextLink(w http.ResponseWriter, pipelineID int, jobName string, page db.Page) {
	w.Header().Add("Link", fmt.Sprintf(
		`<%s/api/v1/pipelines/%d/jobs/%s/builds?%s=%d&%s=%d>; rel="%s"`,
		s.externalURL,
		pipelineID,
		jobName,
		atc.PaginationQueryTo,
		*page.To,
		atc.PaginationQueryLimit,
		page.Limit,
		atc.LinkRelNext,
	))
}

func (s *Server) addPreviousLink(w http.ResponseWriter, pipelineID int, jobName string, page db.Page) {
	w.Header().Add("Link", fmt.Sprintf(
		`<%s/api/v1/pipelines/%d/jobs/%s/builds?%s=%d&%s=%d>; rel="%s"`,
		s.externalURL,
		pipelineID,
		jobName,
		atc.PaginationQueryFrom,
		*page.From,
		atc.PaginationQueryLimit,
		page.Limit,
		atc.LinkRelPrevious,
	))
}
