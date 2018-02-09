import React, { Component } from 'react';
import { connect } from 'react-redux';
import Helmet from 'react-helmet';
import { Link } from 'react-router';
import TalkCards from '../talk-cards';
import { popularTalks } from '../../services';

class Homepage extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    try {
      popularTalks(store)
        .then(() => callback(), callback);
    } catch (err) {
      console.error(err);
      callback();
    }
  }
  /*eslint-enable */

  renderCards({ popularTalksError, popularTalks }) {
    const cards = popularTalks;
    if (popularTalksError || !cards || cards.length === 0) {
      return <div></div>;
    } else {
      return (
        <div>
          <h2 className="title is-2">Popular</h2>
          <TalkCards cards={cards}></TalkCards>
        </div>
      );
    }
  }

  render() {
    return <div>
      <Helmet
        title='Codetalks'
        meta={[{
          property: 'og:title',
          content: 'Track and List your Tech Talk History'
        }]} />
      <h1 className="title is-1"><a href="/">Codetalks</a></h1>
      <p>
        Find, review and keep track of the Tech Talks.
      </p>
      <div className="block">
        <Link className="button is-primary"
          to={'/signin'}>Sign in</Link>&nbsp;
        <Link className="button is-dark"
          to={'/register'}>Register</Link>
      </div>
      {this.renderCards(this.props)}
    </div>;
  }

}

export default connect(store => store)(Homepage);
