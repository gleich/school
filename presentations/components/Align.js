import React from 'react'
import './styles/Align.css'
import 'animate.css'

const Row = ({ children }) => <div className={'row'}>{children}</div>
const Column = ({ position, animation, children }) => {
  if (animation === '' || !animation) {
    switch (position) {
      case 'left':
        animation = 'animate_fadeIn animate__rotateInUpRight'
        break
      case 'right':
        animation = 'animate_fadeIn animate__rotateInUpLeft'
        break
      case 'middle':
        animation = 'animate_fadeInUp'
        break
      default:
        animation = 'animate_fadeIn'
        break
    }
  }
  return <div className={`animate__animated ${animation}`}>{children}</div>
}

export { Row, Column }
