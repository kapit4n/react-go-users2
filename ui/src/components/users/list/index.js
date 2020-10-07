import React from 'react';

export default function Index({ users }) {
  return (<table>
    <tr><td>Name</td></tr>
    {users.map(u => {
      return (<tr key={u.id}>{u.name}</tr>)
    })}
  </table>)
}