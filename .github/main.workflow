workflow "CI" {
  on = "push"
  resolves = ["Go Test"]
}

action "Go Test" {
  uses = "./.github/actions/golang"
  args = "test"
}
