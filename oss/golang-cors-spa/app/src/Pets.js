
import React from 'react'
import Pet from './Pet'

/**
 * Pets component.
 */

export default function({ pets }) {
  if (pets.length) {
    return <ul>
      {pets.map(p => <Pet key={p.name} {...p} />)}
    </ul>
  }

  return <p>Loading pets</p>
}
