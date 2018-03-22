import React, { Component } from 'react';
import { connect } from 'react-redux';
import Helmet from 'react-helmet';
import { Link } from 'react-router';
import TalkCards from '../talk-cards';
import { popularTalks, recentlyAddedTalks } from '../../services';

class Homepage extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    try {
      popularTalks(store)
        .then(() => recentlyAddedTalks(store))
        .then(callback, callback)
        .catch((err) => console.error(err));
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

  renderRecentlyAddedCards({ recentlyAddedTalks, recentlyAddedTalksError }) {
    const cards = recentlyAddedTalks;
    if (recentlyAddedTalksError || !cards || cards.length === 0) {
      return <div></div>;
    } else {
      return (
        <div>
          <h2 className="title is-2">Recently Added</h2>
          <TalkCards cards={cards}></TalkCards>
        </div>
      );
    }
  }

  renderCategories({ categories, err }) {
    categories = categories || [];
    const mapped = categories.map(c => <div>c</div>);
    const displayed = (err || (categories && categories.length === 0)) ?
      (<h4 className="title is-4">Nothing here...</h4>) : mapped;
    return (
      <div>
        <h2 className="title is-2">Categories</h2>
        {displayed}
      </div>
    );
  }

  render() {
    return <div className="page home">
      <Helmet
        title='Codetalks'
        meta={[{
          property: 'og:title',
          content: 'Track and List your Tech Talk History'
        }]} />
      <div className="grid">
        <div>
          <h1 className="title is-1"><a href="/">Codetalks</a></h1>
          <p>
            Learn, Explore and Find Interesting Tech Talks.
          </p>
        </div>
        <div>        
          <div className="block">
            <Link className="button is-primary"
              to={'/signin'}>Sign in</Link>&nbsp;
            <Link className="button is-dark"
              to={'/register'}>Register</Link>
          </div>
        </div>
        {this.renderCards(this.props)}
        {this.renderCategories(this.props)}
        {this.renderRecentlyAddedCards(this.props)}
      </div>
    </div>;
  }

}

export default connect(store => store)(Homepage);
