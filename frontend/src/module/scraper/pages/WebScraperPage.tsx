import { Button, Input } from 'antd';
import React, { useState } from 'react';
import { scraperAPI } from '../../../api/scraper';
import { openNotification } from '../../../meta/globalToaster';
import WebDataComponent from '../component/WebDataComponent';
import { WebSiteDataDto } from '../../../model/scarper';

const WebScraperPage = () => {
  const [loading, setLoading] = useState(false);
  const [siteURL, setSiteURL] = useState('');

  const [response, setResponse] = useState<WebSiteDataDto | null>({
    url: 'https://gihan.orizel.com/',
    title: 'Gihan Gunarathne',
    html_version: 'HTML5',
    heading_levels: { h1: 2, h2: 4, h3: 11, h4: 1, h5: 6, h6: 7 },
    have_login_form: false,
    links: {
      external: [
        'https://www.linkedin.com/in/chamara95/',
        'https://github.com/Chamara-g',
        'https://www.facebook.com/chamara.eng/',
        'https://www.npmjs.com/package/rc-geographic',
        'https://github.com/Chamara-g/rc-geographic',
        'https://packagephobia.com/result?p=rc-geographic',
        'https://npm-stat.com/charts.html?package=rc-geographic',
        'https://snyk.io/test/npm/rc-geographic',
        'https://planner.bueno.com/',
        'https://autorizado.com/',
        'https://test.bueno.com/',
        'https://resortgetaway.com/',
        'https://advisor.resortgetaway.com/',
        'https://ieeexplore.ieee.org/Xplore/home.jsp',
        'https://ieeexplore.ieee.org/document/9185336/',
        'https://medium.com/@chamara95.eng',
        'https://medium.com/@chamara95.eng/neural-network-example-using-fashion-mnist-dataset-c19b48c86cf1',
        'https://medium.com/@chamara95.eng/how-to-install-apache-tomcat-web-server-on-your-laptop-and-run-b83fc4b0424b',
      ],
      inaccessible: [
        'https://test.bueno.com/',
        'https://planner.bueno.com/',
        'https://autorizado.com/',
        'https://www.linkedin.com/in/chamara95/',
        'https://ieeexplore.ieee.org/Xplore/home.jsp',
        'https://ieeexplore.ieee.org/document/9185336/',
        'https://resortgetaway.com/',
        'https://advisor.resortgetaway.com/',
      ],
      internal: ['https://gihan.orizel.com/pdf/gihan_gunarathne.pdf', 'https://gihan.orizel.com/portfolio.html'],
    },
  });

  const onClickScraperButtonClick = () => {
    if (siteURL != '') {
      setResponse(null);
      setLoading(true);

      scraperAPI({ siteURL }).then((data) => {
        try {
          if (data?.status) {
            if (data?.status === 200) {
              successSubmission(data);
            } else {
              failSubmission('Please try again later!');
            }
          } else {
            failSubmission('Please try again later!');
          }
        } catch (error) {
          failSubmission('Please try again later!');
        }
      });
    } else {
      failSubmission('Site URL Required.');
    }
  };

  const successSubmission = (data: any) => {
    setLoading(false);

    setResponse(data);
  };

  const failSubmission = (message: any) => {
    setLoading(false);
    openNotification({ message });
  };

  return (
    <div className="container">
      <div className="row">
        <div className="col-sm-4">
          <div className="row">
            <div className="col-sm-12 mt-3 mb-3">
              <p className="input-box-label">Name</p>
              <Input
                value={siteURL}
                className="input-box"
                placeholder="Enter URL"
                onChange={(e) => {
                  setSiteURL(e.target.value);
                }}
                disabled={loading}
              />
            </div>
            <div className="col-sm-12 mb-3">
              <Button
                type="primary"
                size="large"
                className="submit-button"
                style={{ width: '100%' }}
                onClick={() => {
                  onClickScraperButtonClick();
                }}
                loading={loading}
              >
                Extract Data
              </Button>
            </div>
          </div>
        </div>
        {response && (
          <div className="col-sm-8">
            <WebDataComponent data={response} />
          </div>
        )}
      </div>
    </div>
  );
};

export default WebScraperPage;
