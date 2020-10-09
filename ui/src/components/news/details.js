import React from 'react';

export default function ({ newsInfo }) {
  return (
    <div style={{ width: '30%', display: 'inline-block', border: '1px solid', maxHeight: 400, height: 400, margin: '1rem', overflow: 'hidden' }} >
      <h2>{newsInfo.Title.slice(0, 50)}</h2>
      <div>
        <div>
          Author: {newsInfo.Author}
        </div>
        <img src={newsInfo.URLToImage} width="200" />
        <p>
          {newsInfo.Content}
        </p>
      </div>
    </div>
  )
}