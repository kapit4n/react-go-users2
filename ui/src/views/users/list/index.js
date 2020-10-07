import React from 'react';

import UsersComponent from '../../../components/users/list'
import UserCreateComponent from '../../../components/users/create'
import axios from 'axios';

export default function Index() {

  const [users, setUsers] = React.useState([]);

  React.useEffect(() => {
    async function fetchData() {
      try {
        const res = await axios.get(`${process.env.REACT_APP_API_URL}/users`);
        console.log(res.data.data);
        setUsers(res.data.data);

      } catch (e) {
      }
    }
    fetchData();
  }, [])

  return (
    <>
      <UserCreateComponent />
      <UsersComponent users={users} />
    </>
  )
}