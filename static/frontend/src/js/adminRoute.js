import React from "react";
import { Route, Redirect } from "react-router-dom";

const AdminRoute = ({ component: Component, ...rest }) => (
  <Route {...rest} render={(props) => (
    localStorage.getItem('Jedi') && (document.cookie.indexOf('auth')!==-1)
      ? <Component {...props} />
      : <Redirect to='/challenges' />
  )} />
)

export default AdminRoute;
