package template

const GitlabCI = `job-static-analysis:
  stage: test
  script:
   - make static-analysis

job-check:
  stage: test
  script:
    - make check
`
