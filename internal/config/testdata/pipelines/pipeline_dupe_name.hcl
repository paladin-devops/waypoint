project = "foo"

pipeline "foo" {
  step "test" {
    image_url = "example.com/test"

    use "test" {
      foo = "bar"
    }
  }
}

pipeline "foo" {
  step "test" {
    image_url = "example.com/test"

    use "different_test" {
      foo = "bar"
    }
  }
}

app "web" {
    config {
        env = {
            static = "hello"
        }
    }

    build {}

    deploy {}
}
