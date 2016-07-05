import React from 'react'

import 'styles/searchBar.scss'

export default class SearchBar extends React.Component {
  onKeyDown(e) {
    const { onSearch } = this.props

    if(e.keyCode === 13) {
      onSearch(this.refs.textBox.value)
    }
  }

  render() {
    const { loading, onSearch } = this.props
    return (
      <div className='SearchBar row small-11 medium-6 large-4 small-centered'>
        <input
          className='small-9 columns'
          type='text'
          placeholder='Part of the comic title'
          ref='textBox'
          onKeyDown={(e) => this.onKeyDown(e)}
        />
        <button
          className='button'
          onClick={() => onSearch(this.refs.textBox.value)}
        >
          <i className={`fa ${loading ? 'fa-spinner fa-spin' : 'fa-search'}`}/>
        </button>
      </div>
    )
  }
}
