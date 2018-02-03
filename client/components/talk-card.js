import React, { Component } from 'react';
import { Link } from 'react-router';

export default class TalkCard extends Component {

  static get defaultProps() {
    return {
      card: {
        title: 'Title',
        content: 'Content',
        source: 'Youtube',
        link: 'https://www.youtube.com/watch?v=csyL9EC0S0c',
        linkTitle: '',
      }
    };
  }

  render() {
    const {
      title,
      content,
      source,
      link,
      linkTitle
    } = this.props.card;
    return (
      <div className="card talk-card">
        <p className="card-header-title">
          <a href={link}>{linkTitle}</a>
        </p>
        <div className="card-content">
          <div className="content">
            <p>{content}</p>
            <a>{source} - </a><a href={link}>{linkTitle}</a>
          </div>
        </div>
      </div>
    );
  }
}
