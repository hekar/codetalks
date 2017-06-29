import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router';

export default class Search extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  render() {
    return (
      <div>
        <Helmet title='Codetalks - Search' />
        <h1>Search</h1>
        <input type="text" placeholder="search"/>
        <div className="block">
          <a className="button">Button</a>
        </div>
      </div>
    );
  }

}
