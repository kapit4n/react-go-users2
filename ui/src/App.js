import React from 'react';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
} from 'react-router-dom'

import UsersView from './views/users/list'
import FizzBuzzView from './views/fizz-buzz'

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
              <Link to="/fizzbuzz">Go to Fizz Buzz</Link>
            </li>

          </ul>
        </nav>
        <Switch>

          <Route path="/fizzbuzz">
            <FizzBuzzView />
          </Route>
          <Route path=""></Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
