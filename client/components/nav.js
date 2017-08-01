import React, { Component } from 'react';
import { Link } from 'react-router';

export default class Nav extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    const pathname = this.props;
    return (
      <nav className="nav">
        <div className="nav-left">
          <Link className="nav-item" to={'/'}>
            CodeTalks
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