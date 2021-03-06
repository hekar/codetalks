import React, { Component } from 'react';
import { connect } from 'react-redux';
import Helmet from 'react-helmet';
import { Link } from 'react-router';
import { searchTalks } from '../../services';
import Pagination from '../pagination';
import Grid from '../grid';

class Talks extends Component {
  static onEnter({ store, nextState, replaceState, callback }) {
    try {
      searchTalks(store)
        .then(() => callback(), callback);
    } catch (error) {
      callback();
    }
  }

  render() {
    const { err, talks } = this.props;

    if (err) {
      return <div>Error: {err}</div>;
    } else if (talks) {
      const rows = talks.map(talk =>{
        const tags = talk.tags.map(tag => ((
          <span key={`${talk.id}-tag`} className="tag is-white">
            <a href={'/search/tag/' + tag}>{tag}</a>
          </span>
        )));
        return (
          <div key={talk.id} className="columns list-item">
            <a className="is-2 column" href={talk.url}>
              <img className="thumbnail" src={talk.thumbnailUrl} />
            </a>
            <div className="column">
              <a className="title is-6" href={`/talk/${talk.id}`}>{talk.name}</a>
              <div>{tags}</div>
            </div>
          </div>
        );
      });
      return <div>
        <h1 className="title is-1">Results</h1>
        {rows}
        <Pagination />
      </div>;
    } else {
      return <div>Loading..</div>;
    }
  }

}

export default connect(store => store)(Talks);
