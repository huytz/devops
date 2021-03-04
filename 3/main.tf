
resource "github_repository" "repository" {
  name = "engineering-tools"
}

resource "github_repository" "another_repository" {
  name = "devops-tools"
}

module "team1" {
  source  = "mineiros-io/team/github"
  version = "~> 0.3.0"

  name        = "Engineering"
  description = "This team is created with terraform to test the terraformn-github-repository module."
  privacy     = "closed"

  members     = ["user1,user3,user5,user6,user7,user8,user9,user10"]
  maintainers = ["user2"]

  pull_repositories = [
    github_repository.repository.name,
  ]

  push_repositories = [
    github_repository.another_repository.name,
  ]
}

module "team2" {
  source  = "mineiros-io/team/github"
  version = "~> 0.3.0"

  name        = "Engineering"
  description = "This team is created with terraform to test the terraformn-github-repository module."
  privacy     = "closed"

  members     = ["user2,user3,user4,user6,user7,user8,user9,user10"]
  maintainers = ["user1"]

  pull_repositories = [
    github_repository.repository.name,
  ]

  push_repositories = [
    github_repository.another_repository.name,
  ]
}

module "team3" {
  source  = "mineiros-io/team/github"
  version = "~> 0.3.0"

  name        = "Engineering"
  description = "This team is created with terraform to test the terraformn-github-repository module."
  privacy     = "closed"

  members     = ["user1,user2,user3,user4,user5,user6,user7"]
  maintainers = ["user3"]

  pull_repositories = [
    github_repository.repository.name,
  ]

  push_repositories = [
    github_repository.another_repository.name,
  ]
}

module "team4" {
  source  = "mineiros-io/team/github"
  version = "~> 0.3.0"

  name        = "Engineering"
  description = "This team is created with terraform to test the terraformn-github-repository module."
  privacy     = "closed"

  members     = ["user1,user3,user4,user5,user6,user7,user8,user9,user10"]
  maintainers = ["user2"]

  pull_repositories = [
    github_repository.repository.name,
  ]

  push_repositories = [
    github_repository.another_repository.name,
  ]
}

module "team5" {
  source  = "mineiros-io/team/github"
  version = "~> 0.3.0"

  name        = "Engineering"
  description = "This team is created with terraform to test the terraformn-github-repository module."
  privacy     = "closed"

  members     = ["user2,user3,user4,user6,user7,user8,user9,user10"]
  maintainers = ["user1"]

  pull_repositories = [
    github_repository.repository.name,
  ]

  push_repositories = [
    github_repository.another_repository.name,
  ]
}

module "team6" {
  source  = "mineiros-io/team/github"
  version = "~> 0.3.0"

  name        = "Engineering"
  description = "This team is created with terraform to test the terraformn-github-repository module."
  privacy     = "closed"

  members     = ["user1,user2,user3,user4,user5,user7,user8,user9,user10"]
  maintainers = ["user4"]

  pull_repositories = [
    github_repository.repository.name,
  ]

  push_repositories = [
    github_repository.another_repository.name,
  ]
}