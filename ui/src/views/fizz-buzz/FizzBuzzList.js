import React from 'react';
import axios from 'axios';

export default function FizzBuzzList({ n }) {

  const [fbArray, setFArray] = React.useState([]);

  React.useEffect(() => {
    async function fetchData(n) {
      try {
        if (n > 0) {
          let result = await axios.get(`${process.env.REACT_APP_API_URL}/fizzbuzz/${n}`);
          setFArray(result.data.data.map(element => {
            const parsed = parseInt(element, 10);
            if (isNaN(parsed)) {
              return { element, highlighted: true }
            } else {
              return { element, highlighted: false }
            }
          }));
        }
      } catch (e) {
        console.error(e);
      }
    }

    fetchData(n);

  }, [n]);

  return (<div>
    <table>
      {fbArray.map(f => <tr key={f.element} style={{ color: f.highlighted ? 'green' : 'black' }}><td>{f.element}</td></tr>)}
    </table>
  </div>);
}