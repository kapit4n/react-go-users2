import React from 'react';
import Details from './details';

export default function Index({ newsList }) {
  return (
    <div style={{
      width: '100%'
    }}>
      {newsList.map(newsInfo => <Details key={newsInfo.Title} newsInfo={newsInfo} />)}
    </div >
  )
}
