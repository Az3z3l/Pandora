import React, { Component } from "react";
import  {api} from "../../js/api.js"
import Badge from 'react-bootstrap/Badge'
import "bootstrap/dist/css/bootstrap.css";
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from "reactstrap";
import ReactMarkdown from "react-markdown";

class Challenges extends Component {

	constructor(props) {
		super(props);
		this.state = {
			challenge: [],
			flag: '',
			msg: '',
			msg_type: '',
			openedModal: null,
			auth: true,
			isLoading: true,
			unavailable: false
		}
	}

	openModal = id => {
		this.setState({ openedModal: id });
	};
	closeModal = () => {
		this.setState({ openedModal: null });
	};

	async componentDidMount() {
		var query = `
		query challs{
			challenge {
			  ID
			  name
			  value
			  description
			  category
			  tags
			  solves	
			  flags		  
			  file
			  }
		  }		  
		  `;

		// var datta = await api("challs", {},query)

		let a = await api("challs", {}, query)

		let datum = a.challenge;
		let group=[]
		console.log(datum)
		console.log(datum===null)
		if (datum !== null){
				group = datum.reduce((r, a) => {
				r[a.category[0]] = [...r[a.category[0]] || [], a];
				return r;
			}, {});
		} else {
			this.setState({
				unavailable: true
			})
		}

		this.setState({
			challenge: group
		})


		this.setState({
			isLoading: false
		});

	}

	// openModal = () => this.setState({ isOpen: true });
	// closeModal = () => this.setState({ isOpen: false });


	render() {
		if (this.state.isLoading) {
			return (
				<React.Fragment>
					<div class="spinner"></div>
				</React.Fragment>
			);
		}

		if (this.state.unavailable) {
				return (
					<React.Fragment>
						<div className="container">
							<div className="row mt-5">
								<h2>Challenges</h2>
							</div>
							<br />
							<h3>No Challenge available</h3>
						</div>
					</React.Fragment>
				);
		}
		
		return (
			<React.Fragment>
				<div className="container challenge-font">
					<div className="row">
						<div className="col-md-8">

							{Object.entries(this.state.challenge).map(([key, value]) => {
								return (<>
									<div key={key} className="row mt-5">
										<h3>{key}</h3>
									</div>
									<div className="row">
										{value.map((row) => {
											return (<>
												<div key={row.ID} className="col-sm-4 py-2">
													<div onClick={() => this.openModal(row.ID)} className='card h-100 w-200 bg-dark'>
														<div className="card-body text-center">
															<h6 className="card-title h-50 text-center card">{row.name}</h6>
															<h5 className="card-text h-50 text-center">{row.value}</h5>
															<input type="hidden" name="key" value={row.ID} />
															<div className="challengefont">
																<Modal 
																	isOpen={this.state.openedModal === row.ID}
																	toggle={this.closeModal}
																>
																	<ModalHeader>{row.name}</ModalHeader>
																	<ModalBody>
																		<h5><ReactMarkdown source={row.description} /></h5>
																		<br />
																	solves : {row.solves}
																		<br />
																	categories : <h5>
																			{row.category.map((cats) => {
																				return (<>
																					{' '}<Badge pill variant="danger">
																						{cats}
																					</Badge>
																				</>);
																			}
																			)}</h5>
																		<br />
																	tags : <h5>
																			{row.tags.map((taggy) => {
																				return (<>
																					{' '}<Badge variant="light">
																						{taggy}
																	</Badge>
																				</>);
																			}
																			)} </h5>
																	flag : {row.flags}
																	{(row.file !== null || row.file !== undefined) &&        
																			<h5>
																				{/* File: <a href={`/files/${this.state.challid}/${this.state.challenge.file}`} download>{this.state.challenge.file}</a> */}
																				File: 
																				{row.file.map((cats) => {
																					return (<>
																						{' '}
																						<a href={`/files/${row.ID}/${cats}`} download>{cats}</a>
																					</>);
																				}
																				)}
																			</h5>  
																		}

																	</ModalBody>
																	<ModalFooter>
																		<Button color="primary" onClick={this.closeModal}>
																			Close
																	</Button>
																	</ModalFooter>
																</Modal>
															</div>
														</div>
													</div>
												</div>
											</>);
										})}
									</div>
								</>);
							})}
						</div>
					</div>
				</div>
			</React.Fragment>
		);
	}
}

export default Challenges;
