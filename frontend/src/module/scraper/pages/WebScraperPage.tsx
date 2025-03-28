import { Button, Input } from 'antd';
import React, { useState } from 'react';
import { scraperAPI } from '../../../api/scraper';
import { openNotification } from '../../../meta/globalToaster';
import WebDataComponent from '../component/WebDataComponent';
import { WebSiteDataDto } from '../../../model/scarper';

const WebScraperPage = () => {
  const [loading, setLoading] = useState(false);
  const [siteURL, setSiteURL] = useState('');

  const [response, setResponse] = useState<WebSiteDataDto | null>(null);

  const onClickScraperButtonClick = () => {
    if (siteURL != '') {
      setResponse(null);
      setLoading(true);

      scraperAPI({ siteURL }).then((data) => {
        try {
          if (data?.url) {
            successSubmission(data);
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
