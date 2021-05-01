import './App.css';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
import React, { useState } from 'react';
import NotFound from './components/errors/NotFound';
import Login from './components/auth/Login';
import Register from './components/auth/Register';
import Dashboard from './components/main/Dashboard';
import Profile from './components/main/Profile';
import Leaderboard from './components/main/Leaderboard';
import AuthService from "./services/auth.service";

function App() {
  const [isLogged, setLogged] = useState(AuthService.isLogged());
  return (
    <main>
        <Router>
          <Switch>
            <Route path="/" component={isLogged ? Dashboard : () => <Login setLogged={setLogged} />} exact/>
            <Route path="/register" component={Register} exact />
            <Route path="/dashboard" component={isLogged ? Dashboard : () => <Login setLogged={setLogged} />}/>
            <Route path="/profile" component={isLogged ? Profile : () => <Login setLogged={setLogged} />} />
            <Route path="/leaderboard" component={isLogged ? Leaderboard : () => <Login setLogged={setLogged} />} />
            <Route component={NotFound} />
          </Switch>
        </Router>
    </main>
  )
}

export default App;
