import React, { Component } from "react";
import { api } from  "../js/api.js"
import { ListGroup } from "react-bootstrap"
import Moment from "react-moment";

class Notifications extends Component {
	constructor(props) {
		super(props);
		this.state = {
			notifications: [],
			isLoading: false
		}

	}
	async componentDidMount() {
        var query = `query notifications{
						notify{
						name
						description	
						timestamp
						}
					}`;

		let a = await api("notifications", {}, query)
		
		let res = a.notify
		if (res == null){
			this.setState({
				notifications: []
			});
		} else {
			this.setState({
				notifications: res
			});
		}
		this.setState({
			isLoading: false
		});
	}

	render() {
		if (this.state.isLoading) {
			return (
				<React.Fragment>
					<div class="spinner"></div>
				</React.Fragment>
			);
		}

		return (
			<React.Fragment>
				<div className="container">
					<div className="col-md-12">
						<div className="row mt-5">
							<h2>Notifications</h2>
						</div>
						<div className="row mt-4">
						</div>
						<div className="row">
                        <ListGroup>
						{this.state.notifications.map((row) => {
										return (
                            <ListGroup.Item variant="dark">
                                
								<h4>{row.name}</h4>
								<p>{row.description}</p>
                                <p>	<Moment format="MMM	DD YYYY h:m:s">{row.Timestamp}</Moment></p>

                            </ListGroup.Item>
										)})
						}
						</ListGroup>
						</div>
					</div>
				</div>
			</React.Fragment>
		)
	}
}

export default Notifications;