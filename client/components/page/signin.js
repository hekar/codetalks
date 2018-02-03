import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router';

export default class Signin extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  render() {
    return <div className="content">
      <Helmet title='Codetalks - Sign in' />
      <h1 className="title is=1">Sign In</h1>
      <p>Log into Codetalks to manage your profile, update talks and create lists.</p>
      <form method="POST">
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
          <input
            className="button is-primary"
            type="submit"
            value="sign up"/>
        </div>
      </form>
    </div>;
  }

}
