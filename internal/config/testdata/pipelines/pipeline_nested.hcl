project = "foo"

pipeline "foo" {
  step "test" {
    image_url = "example.com/test"

    use "test" {
      foo = "bar"
    }
  }

  step "nested" {
    pipeline "bar" {
      step "boo" {
        image_url = "example.com/test"

        use "test" {
          foo = "nested"
        }
      }
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
