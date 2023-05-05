package main

func BuscaPlaylistsPorTitulo(titulo string, playlists []playlist) []playlist {
	var playlistsComTitulo []playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == titulo {
				playlistsComTitulo = append(playlistsComTitulo, playlist)
				break
			}
		}
	}

	return playlistsComTitulo
}
