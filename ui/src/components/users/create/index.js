import React from 'react';
import axios from 'axios';

export default function Index() {

  const [userName, setUserName] = React.useState("");
  const [points, setPoints] = React.useState(0);
  const onSubmit = async (e) => {
    e.preventDefault();
    let response = await axios.post(`${process.env.REACT_APP_API_URL}/users`, { name: userName, points })
    console.log(response);
  }

  return (
    <form onSubmit={onSubmit}>
      <div style={{ display: 'flex' }}>
        <span>UserName: </span> <input placeholder="User Name" value={userName} onChange={e => setUserName(e.target.value)} />
        <span>Points: </span> <input placeholder="Points" type='number' value={points} onChange={e => setPoints(e.target.value)} />
        <button type="submit" >Create</button>
      </div>
    </form>
  )
} 