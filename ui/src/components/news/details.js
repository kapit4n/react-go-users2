import React from 'react';

export default function ({ newsInfo }) {
  return (
    <div style={{ width: '30%', display: 'inline-block', border: '1px solid', maxHeight: 400, height: 400, margin: '1rem', overflow: 'hidden' }} >
      <h1>{newsInfo.Title.slice(0, 17)}</h1>
      <div>
        <div>
          Author: {newsInfo.Author}
        </div>
        <img src={newsInfo.URLToImage} width="200" />
        <p>
          {newsInfo.Description}
        </p>
      </div>
    </div>
  )
}