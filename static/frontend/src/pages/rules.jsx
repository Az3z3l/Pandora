import React, { Component } from "react";

class Rules extends Component {
    render() {
        return (
            <React.Fragment>
                <main>
	                <div className="container h-100">
	                    <div className="row align-items-center h-90">
                            <div class="col-sm-12">
	                            <div className="jumbotron bg-transparent">
	                                <h2>The 8 Rules of Archive:</h2>
	                                <hr />
	                                <p>First rule  : Spread the word about the archive.</p>
	                                <p>Second rule : If a challenge seems to be vulnerable in any other way than it was intended, report it to the admins.</p>
	                                <p>Third rule  : This platform is not an attack surface. If you spot any bugs in the platform, report it to the admins.</p>
									<p>Fourth rule : Do not carry out any automated scans on any of the challenges unless otherwise stated.</p>
									<p>Fifth rule  : This archive is meant to help interested people get started in the field of cybersecurity. Thus flag sharing not encouraged.</p>
									<p>Sixth rule  : The flag format might differ from challenge to challenge and will be metioned in the description.</p>
									<p>Seventh rule: Sanity flag is Flag{`{sanity_in_2020?}`}</p>
									<p>Eighth rule : Don't forget to enjoy</p>
									<hr />
									<p>Join the conversation: <a href="https://t.me/joinchat/Jt-vK0yXmdNHzicq8wbhbA" >Telegram</a></p>
	                            </div>
	                        </div>
	                    </div>
	                </div>
                </main>
            </React.Fragment >
        );
    }
}

export default Rules;