import React, { Component } from 'react';
import { Link } from 'react-router-dom';

export default class Nav extends Component {
  render() {
    return (
      <nav className="nav">
        <div className="nav-left">
          <Link className="nav-item" to={'/'}>CodeTalks
          </Link>
          <Link className="nav-item" to={'/talks'}>
            Talks
          </Link>
          <Link className="nav-item" to={'/profile'}>
            Profile
          </Link>
        </div>

        <span className="nav-toggle">
          <span></span>
          <span></span>
          <span></span>
        </span>

        <div className="nav-right nav-menu">
          <label className="label">
            <input name="search" className="input" type="text" placeholder="Search..." />
          </label>
          <a className="button is-black">
            <span className="icon is-small">
              <i className="fa fa-search"></i>
            </span>
          </a>
          &nbsp;
          <Link className="button is-primary"
            to={'/search'}>Sign in</Link>
          &nbsp;
          <Link className="button is-dark"
            to={'/register'}>Register</Link>
        </div>
      </nav>
    );
  }
}