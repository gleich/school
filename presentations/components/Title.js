import React from 'react'
import './styles/Title.css'
import 'animate.css'

export default ({ animation, children }) => (
  <h1
    className={`title animate__animated ${
      animation ? animation : 'animate__fadeInDown'
    }`}
  >
    {children}
  </h1>
)
