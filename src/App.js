import React from 'react';

class App extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			error: null,
			isLoaded: false,
			items: [],
		};
	}

	componentDidMount() {
		fetch('/produce')
			.then((res) => res.json())
			.then(
				(result) => {
					this.setState({
						isLoaded: true,
						items: result,
					});
				},
				(error) => {
					this.setState({
						isLoaded: true,
						error,
					});
				}
			);
	}

	render() {
		const { error, isLoaded, items } = this.state;
		if (error) {
			return <div>Error: {error.message}</div>;
		} else if (!isLoaded) {
			return <div>Loading...</div>;
		} else {
			return (
				<div>
					<h1>Produce List</h1>
					<ul>
						{items.map((item) => (
							<li key={item.id}>{item.name}</li>
						))}
					</ul>
				</div>
			);
		}
	}
}

export default App;
