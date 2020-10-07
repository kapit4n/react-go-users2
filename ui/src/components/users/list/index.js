import React from 'react';

export default function Index({ users }) {
  return (<table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Points</th>
      </tr>
    </thead>
    <tbody>
      {users.map(u => {
        return (
          <tr key={u.id}>
            <td>{u.name}</td>
            <td>{u.points}</td>
          </tr>)
      })}
    </tbody>
  </table>)
}