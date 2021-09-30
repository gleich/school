import React from 'react'
import './styles/Copyright.css'
import 'animate.css'

export default ({ delay }) => (
  <p
    className={`copyright animate__animated animate__slideInUp ${
      delay ? 'animate__delay-' + delay + 's' : ''
    }`}
  >
    &copy; Matt Gleich 2004-{new Date().getFullYear()}{' '}
    <a
      href="https://github.com/gleich/school"
      target="_blank"
      rel="noreferrer"
      className="link"
    >
      gleich/school
    </a>
  </p>
)
