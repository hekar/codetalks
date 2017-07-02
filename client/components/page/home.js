import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router-dom';
import TalkCards from '../talk-cards';

export default class Homepage extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  render() {
    const cards = [{}, {}, {}];
    return <div>
      <Helmet
        title='Codetalks'
        meta={[
          {
            property: 'og:title',
            content: 'Track and List your Tech Talk History'
          }
        ]} />
      <h1 className="title is-1"><a href="/#/">Codetalks</a></h1>
      <p>
        Find, review and keep track of the Tech Talks.
      </p>
      <div className="block">
        <Link className="button is-primary"
          to={'/search'}>Sign in</Link>&nbsp;
        <Link className="button is-dark"
          to={'/register'}>Register</Link>
      </div>
      <div>
        <TalkCards cards={cards}></TalkCards>
      </div>
    </div>;
  }

}
