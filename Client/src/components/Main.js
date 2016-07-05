import React from 'react'
import request from 'superagent'
import 'normalize.css/normalize.css'

import { API_ROOT } from '../config/apiUrl'
import SearchBar from './SearchBar'
import 'styles/app.scss'

class AppComponent extends React.Component {
  constructor() {
    super()
    this.state = {
      comics: [],
      comicIndex: 0,
      loading: false,
      searchResult: []
    }
  }

  fetchComics(index, incComicIndex) {
    if (this.inFetching) {
      return
    }
    this.inFetching = true
    request
      .get(`${API_ROOT}${index}`)
      .end((err, res) => {
        this.inFetching = false
        if (!err) {
          const { comics } = this.state
          const newComics = JSON.parse(res.text)
          if (comics.every(d => d.title != newComics[0].title)) {
            this.setState({ comics: comics.concat(newComics) })
          }
          if (incComicIndex) {
            this.setState({ comicIndex: index })
          }
        } else {
          console.error(err)
        }
      })
  }

  lazyLoadingController() {
    const { offsetHeight: appHeight } = document.querySelector('.App')
    const { scrollTop, offsetHeight } = document.body
    if (appHeight - scrollTop - offsetHeight < 200) {
      const { comicIndex } = this.state
      this.fetchComics(comicIndex + 1, true)
    }
  }

  onSearch(text) {
    if (text === '') {
      return
    }
    this.setState({ loading: true })
    request
      .get(`${API_ROOT}search/${text}`)
      .end((err, res) => {
        this.setState({ loading: false })
        if (!err) {
          let newComics = JSON.parse(res.text)
          newComics = newComics.map(d => {
            d.isSearchRes = true
            return d
          })
          this.setState({ searchResult: newComics })
        } else {
          console.error(err)
        }
      })
  }

  componentWillMount() {
    this.fetchComics(0)
    setInterval(() => {
      this.lazyLoadingController()
    }, 200)
  }

  render() {
    const { comics, loading, searchResult } = this.state
    return (
      <div className='App'>
        <div className='header'>
          <h1 className='title'>xkcdReader</h1>
        </div>
        <p className='description'>Thanks for <a href='https://xkcd.com'>xkcd.com</a> made the comics</p>
        <SearchBar loading={loading} onSearch={text => this.onSearch(text)}/>
        <div className='row small-11 small-centered'>
        {
          searchResult.concat(comics).map((d, key) => (
            <div
              className='small-12 medium-6 large-4 comic-area columns'
              key={key}
            >
              <h3>{d.title}</h3>
              { d.isSearchRes && <p>Came from search</p> }
              <img src={d.img} alt={d.alt} title={d.alt} />
            </div>
          ))
        }
        </div>
      </div>
    )
  }
}

export default AppComponent
