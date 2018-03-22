import React, { Component } from 'react';
import { Link } from 'react-router';
import FontAwesomeIcon from '@fortawesome/react-fontawesome'
import { faSignInAlt, faUserCircle } from '@fortawesome/fontawesome-free-solid'

export default class Nav extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    const pathname = this.props;
    return (
      <nav className="nav bar grid">
        <ul className="links">
          <li><Link className="nav-item" to={'/'}>CodeTalks</Link></li>
          <li><Link className="nav-item" to={'/talks'}>Talks</Link></li>
        </ul>

        <div>
          <form method="GET" action="/talks" className="inline-block">
            <label className="inline-block">
              <input
                name="q"
                className="input"
                type="text"
                placeholder="Search..." />
            </label>
            <input
              className="inline-block"
              type="submit"
              value="Search" />
          </form>
          &nbsp;
          <Link className="inline-block button is-primary"
            to={'/signin'}><FontAwesomeIcon icon={faSignInAlt}/> Sign in</Link>
          &nbsp;
          <Link className="inline-block button is-dark"
            to={'/register'}><FontAwesomeIcon icon={faUserCircle}/> Register</Link>
        </div>
      </nav>
    );
  }
}