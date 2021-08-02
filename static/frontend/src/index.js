import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.css";
import "./css/style.css";
import PrivateRoute from './js/userRoute';
import AdminRoute from "./js/adminRoute"

import { Header, Footer } from "./components/header"
import HomePage from "./pages/HomePage.jsx";
import Login from "./pages/Login.jsx";
import Register from "./pages/Register.jsx";
import Rules from "./pages/rules.jsx";
import ScoreBoard from "./pages/scoreboard.jsx";
import Challenges from "./pages/challenges.jsx";
import Settings from "./pages/settings.jsx";
import Profile from "./pages/profile.jsx";
import Notifications from "./pages/notifications.jsx";

import Adminchallenge from "./pages/admin/challenges.jsx";
import Adminaddchallenge from "./pages/admin/addchallenge.jsx";
import AdminchallengeIndex from "./pages/admin/challengeIndex.jsx";
import AdminChallengeEdit from "./pages/admin/challengeEdit.jsx"
import Adminviewsubmission from "./pages/admin/viewsubmission.jsx";
import Adminaddnotification from "./pages/admin/addnotification.jsx";
import Adminviewuser from "./pages/admin/viewuser.jsx";
import Admindashboard from "./pages/admin/dashboard.jsx";
import Admineditnotification from "./pages/admin/editnotification.jsx";
import Adminindexnotification from "./pages/admin/indexnotification.jsx";
import Adminctfmanagement from "./pages/admin/managementctf.jsx"

import NotFound from "./pages/404.jsx";

ReactDOM.render(
  <Router>
    <Header />
    <Switch>
      <Route exact path={`/`} component={HomePage} />
      <Route path={`/login`} component={Login} />
      <Route path={`/register`}  component={Register} />
      {/* <Route path="/rules" component={Rules} /> */}
      <Route path={`/scoreboard`} component={ScoreBoard} />

      {/* Pages for which the user must be logged in */}
      {/* <PrivateRoute path={`/scoreboard`} component={ScoreBoard} /> */}
      <PrivateRoute path={`/challenges`} component={Challenges} />
      {/* <PrivateRoute path={`/settings`} component={Settings} /> */}
      <PrivateRoute path={`/profile`} component={Profile} />
      <PrivateRoute path={`/notifications`} component={Notifications} />


      {/* routes that might need admin Clearance */}
      {/* <AdminRoute exact path={`/admin/`} component={Admindashboard} /> */}
      <AdminRoute exact path={`/admin/dashboard`} component={Admindashboard} />
      <AdminRoute exact path={`/admin/challenges`} component={Adminchallenge} />
      <AdminRoute exact path={`/admin/challenges/add`} component={Adminaddchallenge} />
      <AdminRoute exact path={`/admin/challenges/edit`} component={AdminchallengeIndex} />
      <AdminRoute exact path={`/admin/challenges/edit/:challid`} component={AdminChallengeEdit} />
      <AdminRoute exact path={`/admin/users/view`} component={Adminviewuser} />
      <AdminRoute exact path={`/admin/users/submissions`} component={Adminviewsubmission} />
      <AdminRoute exact path={`/admin/notifications/add`} component={Adminaddnotification} />
      <AdminRoute exact path={`/admin/notifications/edit/:id`} component={Admineditnotification} />
      <AdminRoute exact path={`/admin/notifications/edit`} component={Adminindexnotification} />
      <AdminRoute exact path={`/admin/ctf/manage`} component={Adminctfmanagement} />
      
      <Route component={NotFound} />
    </Switch>
    <Footer />
  </Router>
  , document.getElementById("root")
)
