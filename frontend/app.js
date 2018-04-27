/* global $:false Handlebars:false */

class Gallery {
  constructor() {
    Handlebars.registerHelper('ifGt', this.handlebarsBiggerThan);
    this.error_tpl = Handlebars.compile($('#error-template').html());
    this.album_tpl = Handlebars.compile($('#album-template').html());

    this.optionset = {
      // General options
      download: false,
      selector: 'img',
      // Thumbnails
      thumbnail: true,
      thumbWidth: 80,
      exThumbImage: 'src',
      showThumbByDefault: false,
    };
  }

  displayError(message = '') {
    $('.container-fluid').append(this.error_tpl({
      message,
    }));
  }

  handlebarsBiggerThan(v1, v2, options) {
    if (v1 > v2) {
      return options.fn(this);
    }
    return options.inverse(this);
  }

  initialize(uri) {
    $.ajax(uri, {
      error: () => this.displayError('Could not fetch albums.'),
      method: 'GET',
      success: (data) => this.processAlbums(data),
    });
  }

  processAlbums(albumData) {
    let albums = albumData.albums;
    albums.sort((a, b) => b.id - a.id);

    albums.forEach((album) => {
      let albumElement = $(this.album_tpl(album));
      $('#gallery-wrap').append(albumElement);
      albumElement.lightGallery(this.updatedOptionSet({
        galleryId: album.id,
      }));
    });
  }

  updatedOptionSet(opts = {}) {
    return Object.assign({}, this.optionset, opts);
  }
}

$(() => {
  let gallery = new Gallery();
  gallery.initialize('albums.json');
});
