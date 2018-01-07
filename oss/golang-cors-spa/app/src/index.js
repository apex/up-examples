
import React from 'react'
import ReactDOM from 'react-dom'
import Pets from './Pets'
import { pets } from './api'

/**
 * Render the app.
 */

function render({ pets }) {
  ReactDOM.render(<Pets pets={pets} />, document.getElementById('app'))
}

// request pets
pets().then(res => {
  render({ pets: res.data })
})

// initial render
render({ pets: [] })
