import React from 'react';

import FizzBuzzList from './FizzBuzzList'

export default function Index() {
  const [n, setN] = React.useState(0);
  const [n2, setN2] = React.useState(0);

  const buildIt = () => {
    setN2(n);
  }

  return (
    <div>
      <input value={n} type="number" onChange={e => setN(e.target.value)} />
      <button onClick={buildIt}> Build</button>
      <FizzBuzzList n={n2} />
    </div>
  )
}