import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router-dom';

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
    </div>;
  }

}
