import React from 'react';

import UsersComponent from '../../../components/users/list'
import axios from 'axios';

export default function Index() {

  const [users, setUsers] = React.useState([]);

  React.useEffect(async () => {
    try{
      const res = await axios.get('http://localhost:8080/users');
      console.log(res.data.data);
      setUsers(res.data.data);

    } catch(e) {

    }
  })

  return <UsersComponent users={users} />
}