package template

const GitlabCI = `job-analyze:
  stage: test
  script:
   - make analyze

job-check:
  stage: test
  script:
    - make check
`