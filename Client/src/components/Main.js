import React from 'react'
import request from 'superagent'
import 'normalize.css/normalize.css'

import { API_ROOT } from '../config/apiUrl'
import 'styles/app.scss'

class AppComponent extends React.Component {
  constructor() {
    super()
    this.state = { comics: []}
  }

  componentWillMount() {
    request
    .get(`${API_ROOT}0`)
    .end((err, res) => {
      if (!err) {
        const { comics } = this.state
        this.setState({ comics: comics.concat(JSON.parse(res.text)) })
      } else {
        console.error(err)
      }
    });
  }

  render() {
    const { comics } = this.state
    return (
      <div className='App'>
        <div className='header'>
          <h1 className='title'>xkcdReader</h1>
        </div>
        <p className='description'>Thanks for <a href='https://xkcd.com'>xkcd.com</a> made the comics</p>
        <div className='row small-11 small-centered'>
        {
          comics.length ? (
            comics.map((d, key) => (
              <div key={key} className='small-12 medium-6 large-4 comic-area columns'>
                <h3>{d.title}</h3>
                <img src={d.img} alt={d.alt}/>
              </div>
            ))
          ) : (
            <p className='indicator'>Loading comics</p>
          )
        }
        </div>
      </div>
    )
  }
}

export default AppComponent
