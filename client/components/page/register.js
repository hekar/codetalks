import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router';

export default class Register extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  render() {
    return <div className="content">
      <Helmet title='Codetalks - Register' />
      <h1 className="title is=1">Register</h1>
      <p>Register your own account on Codetalks.</p>
      <form method="POST">
        <div className="field">
          <label
            htmlFor="username"
            className="label">Username</label>
          <input
            name="username"
            className="input"
            type="text"
            placeholder="username"/>
        </div>
        <div className="field">
          <label
            htmlFor="email"
            className="label">Email</label>
          <input
            name="email"
            className="input"
            type="text"
            placeholder="email"/>
        </div>
        <div className="field">
          <label
            htmlFor="password"
            className="label">Password</label>
          <input
            name="password"
            className="input"
            type="text"
            placeholder="password"/>
        </div>
        <div className="field">
          <label
            htmlFor="password-repeat"
            className="label">Retype Password</label>
          <input
            name="password-repeat"
            className="input"
            type="text"
            placeholder="password (repeat)"/>
        </div>
        <div className="field">
          <p className="control">
            <label className="checkbox">
              <input type="checkbox" />
              I agree to the <a href="#">terms and conditions</a>
            </label>
          </p>
        </div>
        <div className="field">
          <input
            className="button is-primary"
            type="submit"
            value="sign up"/>
        </div>
      </form>
      <div>
        <p>Disclaimer: We won't send you spam.</p>
      </div>
    </div>;
  }

}
