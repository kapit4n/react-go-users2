import React from 'react';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
} from 'react-router-dom'

import UsersView from './views/users/list'
import News from './views/news'

function App() {
  return (
    <Router>
      <div className="App">
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/users">Users</Link>
            </li>

          </ul>
        </nav>
        <Switch>

          <Route path="/users">
            <UsersView />
          </Route>
          <Route path=""><News /></Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
