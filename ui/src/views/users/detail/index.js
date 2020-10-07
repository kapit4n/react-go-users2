import React from 'react';
import axios from 'axios';

import UserComponent from '../../../components/users/detail'

export default function Index({ id }) {

  const [user, setUser] = React.useState();

  // data fetch should go here
  React.useEffect(async () => {
    axios.get('http://localhost:8080/users/' + id).then(user => {
      setUser(user);
    })
  }, [id]);

  return <UserComponent user={user}></UserComponent>;
}