import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import TalkCard from './talk-card';

export default class TalkCards extends Component {
  render() {
    const { cards } = this.props;
    const rendered = cards.map(card => (
      <TalkCard key={card.key}></TalkCard>
    ));
    return (
      <div className="talk-cards">
        {rendered}
      </div>
    );
  }
}