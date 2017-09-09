import React, { Component } from 'react';
import { Link } from 'react-router';
import TalkCard from './talk-card';

export default class TalkCards extends Component {
  render() {
    const { cards } = this.props;
    const rendered = cards
      .map(card => ({
        key: card.id,
        title: card.name,
        tags: card.tags,
        linkTitle: card.name,
        link: card.url,
        thumbnailUrl: card.thumbnailUrl
      }))
      .map(card => (
        <TalkCard key={card.key} card={card}></TalkCard>
      ));
    return (
      <div className="talk-cards">
        {rendered}
      </div>
    );
  }
}