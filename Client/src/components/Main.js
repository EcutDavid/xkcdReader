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
      <div className="index">
        {
          comics.length ? (
            comics.map((d, key) => (
              <div key={key}>
                <h3>{d.title}</h3>
                <img src={d.img} alt={d.alt}/>
                <a
                  href={`http://xkcd.com/${d.num}`}
                  target='_blank'
                >
                  View in xkcd.com
                </a>
              </div>
            ))
          ) : (
            <p className="button">Data fetching</p>
          )
        }
      </div>
    )
  }
}

export default AppComponent
