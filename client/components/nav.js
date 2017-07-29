import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';

export default class Nav extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    const pathname = this.props;
    debugger;
    return (
      <nav className="nav">
        <div className="nav-left">
          <NavLink className="nav-item" to={'/'}>
            CodeTalks
          </NavLink>
          <NavLink className="nav-item" to={'/talks'}>
            Talks
          </NavLink>
          <NavLink className="nav-item" to={'/profile'}>
            Profile
          </NavLink>
        </div>

        <span className="nav-toggle">
          <span></span>
          <span></span>
          <span></span>
        </span>

        <div className="nav-right nav-menu">
          <form method="GET" action="/#/talks">
            <label className="label inline-block">
              <input
                name="q"
                className="input"
                type="text"
                placeholder="Search..." />
            </label>
            <input
              className="inline-block button is-black"
              type="submit"
              value="Search" />
          </form>
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