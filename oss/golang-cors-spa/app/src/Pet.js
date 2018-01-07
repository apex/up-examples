import React from 'react'

/**
 * Pet component.
 */

export default function({ name, species, age }) {
  return <li><strong>{name}</strong> is a {age} year old {species}.</li>
}
