import React, { Component } from 'react';
import { connect } from 'react-redux';
import Helmet from 'react-helmet';
import { Link } from 'react-router';
import { getTalk, getTalkProfile } from '../../services';

class Talk extends Component {
  static onEnter({ store, nextState, replaceState, callback }) {
    const { params } = nextState;
    const { talkId } = params;
    try {
      const promises = [
        getTalk(store, { talkId }),
        getTalkProfile(store, { talkId })
      ];
      Promise.all(promises)
        .then(() => callback(), callback);
    } catch (error) {
      callback();
    }
  }

  render() {
    const {
      errorTalk,
      errorProfile,
      talkLoading,
      profileLoading,
      talk,
      profile,
      meta
    } = this.props.talk;
    if (talkLoading || profileLoading) {
      return <h1>Loading...</h1>;
    }
    if (errorTalk || errorProfile) {
      return <h1>error</h1>
    }
    return (
      <div>
        <h1 className="title is-1">{talk.name}</h1>
        <h4 className="title is-4">Author: {profile.presenter}</h4>
        <p>{profile.summary}</p>
        <code>{JSON.stringify(talk, null, 2)}</code>
        <code>{JSON.stringify(profile, null, 2)}</code>
        <code>{JSON.stringify(meta, null, 2)}</code>
      </div>
    );
  }

}

export default connect(store => store)(Talk);
