package fourOhFour

import (
  "math/rand"
  "time"
)

var Errors = []string {
  "Page not found",
  "Oops! Looks like the page is gone...",
  "This link is broken or the page has moved",
  "Oh no! It looks like the Demogorgean took this page to The Upside Down",
  "Either we broke something or you cannot type",
  "Oops! That page couldn't be found",
  "Oh no, bad luck!",
  "Good job, you broke the internet",
  "Sorry, we couldn't find that page",
  "We couldn't find this page",
  "Congratulations, you broke the internet",
  "We find your lack of navigation disturbing",
  "You underestimate the power of the dark side",
  "Nothing to see here, move along",
  "202 + 202 = ?",
  "Uh oh, something broke",
  "Something has gone horribly wrong!",
  "Stay calm and don't freak out",
}

func Generate() string {
  rand.Seed(time.Now().UnixNano())
  return Errors[rand.Intn(len(Errors) - 1)]
}
