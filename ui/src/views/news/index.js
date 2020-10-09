import React from 'react';
import axios from 'axios';
import NewsComponent from '../../components/news'

export default function Index() {

  const [news, setNews] = React.useState([]);

  React.useEffect(() => {

    async function load() {
      const url = 'http://localhost:8080/news/query';
      const response = await axios.get(url);
      console.log(response);
      setNews(response.data.Articles);
    }

    load();
  }, [])

  return <NewsComponent newsList={news} />
}