import './App.css';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"

import NotFound from './components/errors/NotFound';
import Login from './components/auth/Login';
import Register from './components/auth/Register';

function App() {
  return (
    <main>
      <Router>
        <Switch>
          <Route path="/" component={Login} exact />
          <Route path="/register" component={Register} />
          <Route component={NotFound} />
        </Switch>
      </Router>
    </main>
  )
}

function Home() {
  return (
    <div><h3>home</h3>
    </div>
  )
}

export default App;
